# Go Simple JWT

Code dibawah ini ditulis dengan bahasa pemrograman go dengan library chi

## Gimana cara jalaninnya

```bash
go run main.go
```

```bash
curl http://localhost:3001
```

Response:

proof of concept securing with jwt - public api, token: eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VyX2lkIjoxMjN9.PZLMJBT9OIVG2qgp9hQr685oVYFgRgWpcSPmNcw6y7M

```bash
curl -H "Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VyX2lkIjoxMjN9.PZLMJBT9OIVG2qgp9hQr685oVYFgRgWpcSPmNcw6y7M" http://localhost:3001/profile
```

Response:

protected area. hi 123

```bash
curl http://localhost:3001/profile
```

Response:

Unauthorized