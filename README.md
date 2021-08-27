# Получить инфу о файле

https://api.telegram.org/bot{bot-token}/getFile?file_id={file-id}

```json
{
    "ok": true,
    "result": {
        "file_id": "BQACAgIAAxkBAAMFYR5Wo_JQlqCFHDS_4-IpPRKWg4IAAuYLAAL3WfhIhX2E1ieLDyMgBA",
        "file_unique_id": "AgAD5gsAAvdZ-Eg",
        "file_size": 1860686,
        "file_path": "documents/file_0.PNG"
    }
}
```

# Получить файл

https://api.telegram.org/file/bot{bot-token}/documents/{file_path}

```json
  --> FILE
```