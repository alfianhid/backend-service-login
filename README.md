# Login to User's Account

Login to an account for the non-authenticated user if an account for that user has already exist.

**URL** : `/api/login`

**Method** : `POST`

**Auth required** : NO

**Permissions required** : None

**Data constraints**

Provide the data of user's account to be created.

```json
{
  "username": "[required,min=3,alphanumeric]",
  "password": "[required,min=8]"
}
```

**Data example**

```json
{
  "username": "ahmadakbar",
  "password": "password"
}
```

## Success Response

**Condition** : If everything is OK and an account is exist for the user.

**Code** : `200 OK`

**Content example**

```json
{
  "status": 200,
  "message": "2c7b55a6-d513-4682-808f-9d589b5a100a",
  "accessToken": "eyJhbGciOiJSUzI1NiIsInR5cCIgOiAiSldUIiwia2lkIiA6ICJFTnNBbXZYRnRSN1NlM1B0T2ExSmtVTTUxZWc1cW1EX1ZkSHhUS202Z3hvIn0.eyJleHAiOjE2NDQ1NjkzNTgsImlhdCI6MTY0NDU2OTA1OCwianRpIjoiMzVkMWE2ZjItY2U0NS00NTRkLWFhMjEtZDg1NDM1ZTVlMjBhIiwiaXNzIjoiaHR0cDovL2xvY2FsaG9zdDo4MTgwL2F1dGgvcmVhbG1zL2VkbSIsImF1ZCI6ImFjY291bnQiLCJzdWIiOiIyYzdiNTVhNi1kNTEzLTQ2ODItODA4Zi05ZDU4OWI1YTEwMGEiLCJ0eXAiOiJCZWFyZXIiLCJhenAiOiJlZG1wb3J0YWwiLCJzZXNzaW9uX3N0YXRlIjoiZWM3YTJkMDctNzBhMi00MTUxLWEzZmMtNWRhOWIzYmVmZTBiIiwiYWNyIjoiMSIsInJlYWxtX2FjY2VzcyI6eyJyb2xlcyI6WyJvZmZsaW5lX2FjY2VzcyIsImRlZmF1bHQtcm9sZXMtZWRtIiwidW1hX2F1dGhvcml6YXRpb24iXX0sInJlc291cmNlX2FjY2VzcyI6eyJhY2NvdW50Ijp7InJvbGVzIjpbIm1hbmFnZS1hY2NvdW50IiwibWFuYWdlLWFjY291bnQtbGlua3MiLCJ2aWV3LXByb2ZpbGUiXX19LCJzY29wZSI6ImVtYWlsIHByb2ZpbGUiLCJzaWQiOiJlYzdhMmQwNy03MGEyLTQxNTEtYTNmYy01ZGE5YjNiZWZlMGIiLCJlbWFpbF92ZXJpZmllZCI6ZmFsc2UsIm5hbWUiOiJBaG1hZCBBa2JhciIsInByZWZlcnJlZF91c2VybmFtZSI6ImFobWFkYWtiYXIiLCJnaXZlbl9uYW1lIjoiQWhtYWQiLCJmYW1pbHlfbmFtZSI6IkFrYmFyIiwiZW1haWwiOiJhaG1hZGFrYmFyQGdtYWlsLmNvbSJ9.Z5NhfWFk_JFQ8bUSNnQvnYvrMA3eEEED5hqx4VKquQ_If9a1b00cSJtVGCPdPp0Flxyn0AyjnvyW-i6Og01ObcGxHX0JAfh-nb0uNwbBmbX6e2KgNl1VIFA8ZciX0XT5jiD9lVtBSi0Rv07qUzwxQg9P1nfwAGWU1BJmfz3YYOGJlfYv0JvQOKNHzCcx7D1OYIw0JLGRS7gScr6RGJjQb5i4yGX1BdULWgoEnDi4MYEtF2rh6WeRgzxQpm8ROxmbzu5F1zJwulIn79vQuc4neTVSHtohO-l7_csjX7_XDJZjuGxe7uk5VbkHbHgooAe_Wp25x4k8F3y0cmeO0y6aVg",
  "refreshToken": "eyJhbGciOiJIUzI1NiIsInR5cCIgOiAiSldUIiwia2lkIiA6ICJkMDc4MjhmMi0yN2VmLTRmNGYtYTgwMy03NjFlYjRkNTdmMDUifQ.eyJleHAiOjE2NDQ1NzA4NTgsImlhdCI6MTY0NDU2OTA1OCwianRpIjoiMzA0YThlNTgtNzE4OC00ZjllLWIxYjQtOGFhYzcyODRiMWFmIiwiaXNzIjoiaHR0cDovL2xvY2FsaG9zdDo4MTgwL2F1dGgvcmVhbG1zL2VkbSIsImF1ZCI6Imh0dHA6Ly9sb2NhbGhvc3Q6ODE4MC9hdXRoL3JlYWxtcy9lZG0iLCJzdWIiOiIyYzdiNTVhNi1kNTEzLTQ2ODItODA4Zi05ZDU4OWI1YTEwMGEiLCJ0eXAiOiJSZWZyZXNoIiwiYXpwIjoiZWRtcG9ydGFsIiwic2Vzc2lvbl9zdGF0ZSI6ImVjN2EyZDA3LTcwYTItNDE1MS1hM2ZjLTVkYTliM2JlZmUwYiIsInNjb3BlIjoiZW1haWwgcHJvZmlsZSIsInNpZCI6ImVjN2EyZDA3LTcwYTItNDE1MS1hM2ZjLTVkYTliM2JlZmUwYiJ9.Ut5KiXJwwrI90UTVpaYYvI-1MaIvfSqJfJt1ZsB82c8",
  "expiresIn": 300
}
```

## Error Responses

**Condition** : If user's account credential is wrong for the user.

**Code** : `400 BAD REQUEST`

**Content example**

```json
"code=404, message=username not found"
```

```json
"code=400, message=wrong password"
```

### and

**Condition** : If fields are missed or invalid format.

**Code** : `400 BAD REQUEST`

**Content example**

```json
"code=400, message=payload is unknown"
```
