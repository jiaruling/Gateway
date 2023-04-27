# 产生假数据

from time import sleep
import requests
from random import choice, randint
import ssl

ssl._create_default_https_context = ssl._create_unverified_context

url = (
    "http://127.0.0.1:8080/test_strip_uri/abc",
    "http://127.0.0.1:8080/test_strip_uri/aaa",
    "http://127.0.0.1:8080/test_strip_uri/ccc",
    "http://127.0.0.1:8080/test_strip_uri/bbb",
    "http://127.0.0.1:8080/test_strip_uri/cba",
    # "https://127.0.0.1:4433/test_https_server/bba"
)

payload = {}
headers = {
    'Authorization': 'Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE2ODI1Nzk3NjYsImlzcyI6ImFwcF9pZF9hIn0.oSNI6MBX9O6rbNOXj16uM3IBXxDhJ9EPUPHwLmgGGM8'
}

for i in range(0, 100):
    response = requests.get(choice(url), headers=headers, data=payload)
    if response.status_code == 200:
        print('Request succeeded!')
        print(response.text)
    else:
        print('Request failed with status code', response.status_code)
