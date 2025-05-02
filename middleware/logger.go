
package middleware

import (
    "github.com/gin-gonic/gin"
    "github.com/sirupsen/logrus"
    "os"
    "time"
)

func Logger() gin.HandlerFunc {
    log := logrus.New()
    log.SetFormatter(&logrus.JSONFormatter{})

    appFile, _ := os.OpenFile("logs/app.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)
    errFile, _ := os.OpenFile("logs/error.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)

    log.SetOutput(appFile)
    log.AddHook(&levelHook{Writer: errFile, LogLevels: []logrus.Level{logrus.WarnLevel, logrus.ErrorLevel, logrus.FatalLevel}})
    return func(c *gin.Context) {
        start := time.Now()
        c.Next()
        latency := time.Since(start)
        status := c.Writer.Status()
        entry := log.WithFields(logrus.Fields{
            "status":  status,
            "method":  c.Request.Method,
            "path":    c.Request.URL.Path,
            "latency": latency,
            "client":  c.ClientIP(),
        })
        if len(c.Errors) > 0 {
            entry.Error(c.Errors.String())
        } else {
            entry.Info("request completed")
        }
    }
}

type levelHook struct {
    Writer    *os.File
    LogLevels []logrus.Level
}

func (h *levelHook) Fire(entry *logrus.Entry) error {
    line, _ := entry.String()
    _, err := h.Writer.Write([]byte(line))
    return err
}

func (h *levelHook) Levels() []logrus.Level {
    return h.LogLevels
}
