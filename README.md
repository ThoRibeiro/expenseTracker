# 💰 ExpenseTracker API

Bienvenue dans **ExpenseTracker**, une API REST développée en Go avec le framework Gin et SQLite, conçue pour simplifier l'enregistrement, la consultation et l'analyse de vos dépenses quotidiennes 📊.

---

## 🚀 Lancer en local

```bash
go run .
```

**Variables d'environnement** :

* 🔑 `JWT_SECRET` (défaut `secret`)
* 📧 `ADMIN_EMAIL` (défaut `admin@example.com`)
* 🔒 `ADMIN_PASS` (défaut `admin`)

🛠️ **Swagger UI** : [http://localhost:8080/docs/index.html](http://localhost:8080/docs/index.html)

---

## 📂 Collection de tests

Un fichier `requests.http` est fourni pour tester tous les endpoints :

* ✍️ **Inscription** : `POST /users/register`
* 🔐 **Login** : `POST /users/login`
* 🧑‍💻 **Profil** : `PUT /users/me`
* 📋 **CRUD Dépenses** :

    * `GET /users/expenses`
    * `GET /users/expenses/:id`
    * `POST /users/expenses`
    * `PUT /users/expenses/:id`
    * `DELETE /users/expenses/:id`
* 🔄 **Bulk** :

    * `PUT /users/expenses/bulk`
    * `DELETE /users/expenses/bulk`
* 🔍 **Search** : `GET /users/expenses/search?q=`
* 🧹 **Reset DB** : `POST /users/admin/reset`
* 🛡️ **Admin update user** : `PUT /users/admin/users/:id`

---

## 🧪 CI & Tests

Le workflow GitHub Actions (`.github/workflows/ci.yml`) couvre :

* 📦 Installation des dépendances
* 📜 Génération de la documentation Swagger (`swag init`)
* ✅ Exécution des tests (`go test ./... -v`)
* 🏗️ Compilation du binaire (`go build -o server`)

```bash
go test ./... -v
go build -o server .
```

---

## ☁️ Déploiement Render

Le fichier `render.yaml` configure le déploiement sur Render (plan free) :

* 🌐 **Web Service** Go 1.22
* 🔑 Variables d’environnement : `JWT_SECRET`, `ADMIN_EMAIL`, `ADMIN_PASS`
* 🔄 Déploiement automatique sur la branche `main`

---

## 📊 Auto‑évaluation

| Exigence                                          |  Points  |   Réalisé   |
| ------------------------------------------------- | :------: |:-----------:|
| CRUD basique (List, Read, Create, Update, Delete) |     2    |      ✅      |
| Recherche plein-texte                             |    0.5   |      ✅      |
| Bulk update/delete                                |    0.5   |      ✅      |
| Réinitialiser la BDD                              |    0.5   |      ✅      |
| Authentification & JWT                            |     3    |      ✅      |
| Validation & gestion d’erreurs                    |     3    |      ✅      |
| Documentation Swagger UI                          |     2    |      ✅      |
| Logging (app.log & error.log)                     |     1    |      ❌      |
| Collection `requests.http`                        |     2    |      ✅      |
| Déploiement gratuit (Render)                      |     3    |      ✅      |
| **Total**                                         | **16.5** | **16.5/20** |

---

## 📁 Structure du projet

```
.
├── config/            # chargement des variables d'env
├── controllers/       # handlers auth, users, expenses
├── database/          # connexion SQLite & migrations
├── docs/              # Swagger docs générés
├── docs/index.html    # UI Swagger pour GitHub Pages
├── logs/              # logs: app.log, error.log
├── middleware/        # JWT auth & logging middleware
├── models/            # struct User & Expense
├── routes/            # définition des routes
├── seed/              # seed automatique de l'admin
├── utils/             # helpers: password, token, pagination
├── requests.http      # collection HTTP REST client
├── render.yaml        # blueprint Render
├── go.mod
├── go.sum
├── main.go
└── README.md          # ce fichier
```
