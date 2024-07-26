import pyautogui
import pyperclip
import time
import csv
import pandas as pd
import os
import psycopg2

conn_params = {
    'dbname': 'defaultdb',
    'user': 'doadmin',
    'password': 'AVNS_rlB4fN4fccJEbMilccC',
    'host': 'oggdb-do-user-16975301-0.c.db.ondigitalocean.com',  # or your host
    'port': '25060'        # default port
}

conn = psycopg2.connect(**conn_params)

query = "SELECT * FROM lh_viber where status = 0;"

df = pd.read_sql_query(query, conn)

conn.close()

for i in range(len(df)):
    print(df.iloc[i, 1])

# pyautogui.moveTo(100, 40, duration = 1)
# pyautogui.click(100,40)

# pyautogui.moveTo(280, 80, duration = 1)
# pyautogui.click(280,80)

# pyautogui.typewrite("88029551")


# pyautogui.moveTo(100, 220, duration = 1)
# pyautogui.click(100,220)

# time.sleep(5)
# text_to_type = "Сайн байна уу, сайхан наадаж байна уу"
# pyperclip.copy(text_to_type)
# pyautogui.hotkey('command', 'v')
# pyautogui.typewrite(["enter"])



# pyautogui.moveTo(320, 855, duration = 1)
# pyautogui.click(320,855)

# pyautogui.moveTo(350, 350, duration = 1)
# pyautogui.click(350,350)


# pyautogui.moveTo(550, 320, duration = 1)
# pyautogui.click(550,320)

# text_to_type = "Танилцуулга"
# pyperclip.copy(text_to_type)
# pyautogui.hotkey('command', 'v')
# pyautogui.typewrite(["enter"])
