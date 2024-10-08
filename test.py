import requests

data = {
    "key": "test",
    "value": "this is a value"
}

r = requests.post("http://127.0.0.1:8000/create", json=data)

print(r.text)
