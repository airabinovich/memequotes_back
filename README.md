# Memequotes

This is a dumb application to use as template for a Golang + Gin-gonic backend app

## Endpoints

### POST /character
Creates a new Character. The body for the call should be
```json
{
  "name": "character_name"
}
```

### GET /character/:id
Retrieve the Character matching the Id. The response body should be
```json
{
  "id": 1,
  "name": "character_name",
  "date_created": "2020-06-14T17:45:00.000Z",
  "last_updated": "2020-06-14T17:45:00.000Z"
}
```

### PATCH /character/:id
Edit a Character. The body should be
```json
{
  "name": "new_character_name"
}
```

### DELETE /character/:id
Delete a character. No body for response, status 410 if deleted

### GET /character/:id/phrases
Retrieve all phrases from a character. Response body:
```json
{
  "results": [
    {
      "id": 1,
      "name": "phrase name",
      "content": "phrase content",
      "date_created": "2020-06-14T17:45:00.000Z",
      "last_updated": "2020-06-14T17:45:00.000Z"
    }
  ]
}
```

### POST /character/:id/phrase
Create a new phrase for a character. The body:
```json
{
  "name": "phrase name",
  "content": "phrase content"
}
```