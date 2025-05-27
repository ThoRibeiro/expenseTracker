package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"os"
	"time"
)

func Logger() gin.HandlerFunc {
	// Création du dossier logs
	if err := os.MkdirAll("logs", 0755); err != nil {
		panic("Impossible de créer logs/: " + err.Error())
	}

	log := logrus.New()
	log.SetFormatter(&logrus.JSONFormatter{})

	// Ouverture app.log
	appFile, err := os.OpenFile("logs/app.log",
		os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)
	if err != nil {
		log.Warn("Impossible d’ouvrir app.log, stdout : " + err.Error())
	} else {
		log.SetOutput(appFile)
	}

	// Ouverture error.log et hook
	errFile, err := os.OpenFile("logs/error.log",
		os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)
	if err != nil {
		log.Warn("Impossible d’ouvrir error.log, logs d’erreur vers stdout : " + err.Error())
	} else {
		log.AddHook(&levelHook{
			Writer:    errFile,
			LogLevels: []logrus.Level{logrus.WarnLevel, logrus.ErrorLevel, logrus.FatalLevel},
		})
	}

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
			for _, e := range c.Errors.Errors() {
				entry.Error(e)
			}
		} else {
			entry.Info("request completed")
		}
	}
}

// levelHook écrit les entrées warn/error/fatal dans error.log
type levelHook struct {
	Writer    *os.File
	LogLevels []logrus.Level
}

// Fire est appelé pour chaque log de niveau Level()
func (h *levelHook) Fire(entry *logrus.Entry) error {
	line, _ := entry.String()
	_, err := h.Writer.Write([]byte(line))
	return err
}

// Levels indique sur quels niveaux ce hook doit être déclenché
func (h *levelHook) Levels() []logrus.Level {
	return h.LogLevels
}
