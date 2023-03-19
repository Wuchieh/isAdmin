import os
from ctypes import *

# ptr = WinDLL(".\\isAdmin.dll")
ptr = cdll.LoadLibrary(".\\isAdmin.dll")

if ptr.isAdmin():
    print("Hi Administrator")
else:
    print("Hi Friend")

os.system("PAUSE")
