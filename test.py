import requests

data = {
    "Key": "nigga",
    "Value": "this is a value"
}

r = requests.post("http://127.0.0.1:8000/create", json=data)

print(r.text)
