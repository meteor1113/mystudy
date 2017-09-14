#-*- coding:utf-8 -*-

import time
import datetime
import random
import requests
import wxpy


KEY = '88b517bdc02a4c28a9ae4a6854da4703'

def get_response(msg):
    Url = 'http://www.tuling123.com/openapi/api'

    data = {
        'key'  : KEY,
        'info'  : msg,
        'userid' : 'pth_robot',
    }

    try:
        r = requests.post(Url, data=data).json()
        return r.get('text')
    except Exception as e:
        print e
        return u"我迷糊了……"


bot = wxpy.Bot(cache_path=True, console_qr=2)
#bot.file_helper.send("login")

@bot.register(wxpy.Group, wxpy.TEXT)
def auto_reply(msg):
    try:
        timeout = random.randrange(0, 50)
        #print "Recv[%s@%s](%s): %s" % (msg.member.name, msg.sender.name, msg.text, timeout)
        print u"{} Recv[{}@{}]({:>2}): {}".format(datetime.datetime.now().isoformat(), msg.member.name, msg.sender.name, timeout, msg.text)
        if msg.is_at or timeout < 5 or timeout > 45 or msg.text.find(u"妙妙") >= 0 or msg.text.find(u"聊天") >= 0 or msg.text.find(u"机器人") >= 0:
            #return "Recv: " + msg.text
            #print "wait %s seconds ..." % (timeout / 10)
            time.sleep(timeout / 10)
            reply = u"睡觉了，别吵妙妙好吗，Zzz……" if timeout < 5 else get_response(msg.text)
            reply = "@" + msg.member.name + " " + reply if msg.is_at else reply
            print u"{} Send: {}".format(datetime.datetime.now().isoformat(), reply)
            return reply
        elif msg.text.find(u'奇奇') >= 0 or msg.text.find(u'奇妙') >= 0:
            #print "wait %s seconds ..." % (timeout / 10)
            time.sleep(timeout / 10)
            reply = u"奇奇是我姐姐，我们是奇妙姐妹!"
            print u"{} Send: {}".format(datetime.datetime.now().isoformat(), reply)
            return reply
        else:
            return
    except Exception as e:
        print e
        return

wxpy.embed()
