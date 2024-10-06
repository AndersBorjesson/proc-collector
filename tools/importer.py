#!/usr/bin/env python3
# -*- coding: utf-8 -*-
"""
Created on Tue Oct  1 21:06:11 2024

@author: aborjes3
"""



from os import listdir
from os.path import isfile, join
import json


mypath = "../src/"

def listFiles(path):
    onlyfiles = [f for f in listdir(mypath) if isfile(join(mypath, f))]
    jsons = []
    for l1 in onlyfiles:
        if ".memdump.json" in l1:
            jsons.append(l1)
    return sorted(jsons)


def readJson(path, jsons):
    data = []
    for l1 in jsons:
        print(l1)
        p = join(path, l1)
        with open(p, 'r') as f:
            d = json.load(f)
            data.append(d)
    return data

j = listFiles(mypath)

data = readJson(mypath, j)