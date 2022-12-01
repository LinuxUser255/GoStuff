#!/usr/bin/env python3

import os

def cat_go_imports():
    
    cmd = 'cat '
    filename = 'i.txt'
    with open(filename, 'r') as f:
        os.system(cmd + ' ' + filename) 

    
cat_go_imports()
