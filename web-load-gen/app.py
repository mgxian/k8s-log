import os
import time
import random
import logging
import requests

URL = ['/', '/a', '/b', '/c', '/d']
USER_AGENT = ['pc', 'mobile', 'app']
INTERVAL = os.getenv('INTERVAL')
ROOT_URL = os.getenv('ROOT_URL', default="http://127.0.0.1")

try:
    INTERVAL = int(INTERVAL)
except:
    INTERVAL = 1

def gen_load():
    while True:
        try:
            headers = {
                'User-Agent': random.choice(USER_AGENT)
            }
            url = ROOT_URL + random.choice(URL)
            requests.get(url, headers=headers)
            logging.info('get success url: {}'.format(url))
            time.sleep(INTERVAL)
        except Exception as e:
            logging.error('get error url: {}'.format(url))
            logging.error(e)


if __name__ == "__main__":
    gen_load()
