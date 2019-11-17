import time

from wxpy import Bot
import json
import requests
from urllib.parse import urlencode


class JsonDb:

    def __init__(self, path):
        self.path = path
        self.data = self.read_config_from_json()

        self.name = self.data["name"]
        self.last_chapter_name = self.data['last_chapter_name']

    def read_config_from_json(self):
        return json.load(open(self.path))

    def save_to_json(self):
        with open(self.path, 'w') as f:
            self.data['name'] = self.name
            self.data['last_chapter_name'] = self.last_chapter_name
            json.dump(self.data, f, ensure_ascii=False)


class BookPushBot:
    CHAPTER_URL = "http://127.0.0.1:8080/api"
    PUSH_URL = "http://127.0.0.1:8080/render"

    def __init__(self, db):
        self.db = db
        self.bot = None
        self.push_url = self.PUSH_URL + "?" + urlencode({"name": self.db.name})

    def login(self):
        print('登录微信中...请扫描二维码')
        self.bot = Bot(cache_path=True)

    def get_last_chapter_name(self):
        r = requests.get(self.CHAPTER_URL, params={"name": self.db.name})
        return r.json()['LastChapterName']

    def push_to_ireader(self):
        ireader = self.bot.search('掌阅iReader')[0]
        ireader.send(self.push_url)
        print(f"push to ireader...")

    def run(self):
        print(f"当前追书: {db.name} 已经追到的章节: {db.last_chapter_name}")
        self.login()
        self.cron_job()

    def cron_job(self):
        while True:
            now_last_chapter = self.get_last_chapter_name()
            print(f"目前网上最新章节: {now_last_chapter}")
            if now_last_chapter != self.db.last_chapter_name:
                self.push_to_ireader()
                self.db.last_chapter_name = now_last_chapter
                self.db.save_to_json()
            else:
                print("你已经看到最新章节了....休眠中")
            time.sleep(60*60)


if __name__ == "__main__":
    db = JsonDb("config.json")
    bot = BookPushBot(db)
    bot.run()
