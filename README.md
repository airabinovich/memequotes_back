# Memequotes

This is a dumb application to use as template for a Golang + Gin-gonic backend app

## Database

### Structure

The database structure is in `db_structure.sql`

### Host and Credentials

The DB host and credentials should be in a file `credentials.conf` (added in .gitignore) with format

```conf
db.host=localhost
db.port=3306
db.user=root
db.password=password
```

Run the application using
```sh
go run main.go --credentials=credentials.conf
```

## Endpoints

### POST /character
Creates a new Character. The body for the call should be
```json
{
  "name": "character_name"
}
```

### GET /characters
Retrieve all Characters. The response body should be
```json
{
  "results": [
    {
      "id": 1,
      "name": "character_name",
      "date_created": "2020-06-14T17:45:00.000Z",
      "last_updated": "2020-06-14T17:45:00.000Z"
    }
  ]
}
```

### GET /character/:character-id
Retrieve the Character matching the Id. The response body should be
```json
{
  "id": 1,
  "name": "character_name",
  "date_created": "2020-06-14T17:45:00.000Z",
  "last_updated": "2020-06-14T17:45:00.000Z"
}
```

### PATCH /character/:character-id
Edit a Character. The body should be
```json
{
  "name": "new_character_name"
}
```

### DELETE /character/:character-id
Delete a character. No body for response, status 410 if deleted

### GET /character/:character-id/phrase/:phrase-id
Retrieve a phrases from a character, only if it belongs to that character. Response body:
```json
{
  "id": 1,
  "content": "phrase content",
  "date_created": "2020-06-14T17:45:00.000Z",
  "last_updated": "2020-06-14T17:45:00.000Z"
}
```

### GET /character/:character-id/phrases
Retrieve all phrases from a character. Response body:
```json
{
  "results": [
    {
      "id": 1,
      "content": "phrase content",
      "date_created": "2020-06-14T17:45:00.000Z",
      "last_updated": "2020-06-14T17:45:00.000Z"
    }
  ]
}
```

### POST /character/:character-id/phrase
Create a new phrase for a character. The body:
```json
{
  "content": "phrase content"
}
```

### DELETE /character/:character-id/phrase/:phrase-id
Delete a phrase matching the phrase-id, only if it belongs to the character-id. No body for response, status 410 if deleted
