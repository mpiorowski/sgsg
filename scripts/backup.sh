#!/bin/bash

sqlite3 ../server/system/db.sqlite3 ".backup ../server/system/db.sqlite3.bak"
