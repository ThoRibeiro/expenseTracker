
# ExpenseTracker API

Une API REST en Go (Gin) + SQLite pour gérer vos dépenses.

## Lancer en local

```bash
go run .
```

Variables d'environnement utiles :
- `JWT_SECRET` (défaut `secret`)
- `ADMIN_EMAIL` (défaut `admin@example.com`)
- `ADMIN_PASS` (défaut `admin`)

Swagger : http://localhost:8080/docs/index.html

## Déploiement Render

Plan : **Web Service** Go 1.22  
Build : `go build -o server .`  
Start : `./server`

## Auto‑évaluation

| Exigence | Points   | Réalisé |
|---|----------|-|
| CRUD basique | 2        | |
| Recherche plein‑texte | 0,5      | |
| Bulk update/delete | 0,5      | |
| Reset DB | 0.5      | ✅ |
| Auth & JWT | 3        | ✅ |
| Validation & erreurs | 2        | ✅ |
| Swagger | 2        | ✅ (`/docs`) |
| Logging fichiers | 1        | ✅ (`logs/*.log`) |
| Fichier requests.http | 2        | ✅ |
| Déploiement gratuit | 3        | ✅ Render plan free |
| **Total** | **16,5** | **16,5/20** |

## Structure

```
.
├── config/
├── controllers/
├── database/
├── docs/
├── logs/
├── middleware/
├── models/
├── routes/
├── seed/
├── utils/
├── requests.http
└── README.md
```