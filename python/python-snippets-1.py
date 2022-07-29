## Merge Two Lists Into a Dictionary

keys_list = ['A', 'B', 'C']
values_list = ['blue', 'red', 'bold']

#There are 3 ways to convert these two lists into a dictionary
#1- Using Python's zip, dict functionz
dict_method_1 = dict(zip(keys_list, values_list))

#2- Using the zip function with dictionary comprehensions
dict_method_2 = {key:value for key, value in zip(keys_list, values_list)}

#3- Using the zip function with a loop
items_tuples = zip(keys_list, values_list) 
dict_method_3 = {} 
for key, value in items_tuples: 
    if key in dict_method_3: 
        pass # To avoid repeating keys.
    else: 
        dict_method_3[key] = value

## Merge Two or More Lists Into a List of Lists

def merge(*args, missing_val = None):
#missing_val will be used when one of the smaller lists is shorter tham the others.
#Get the maximum length within the smaller lists.
  max_length = max([len(lst) for lst in args])
  outList = []
  for i in range(max_length):
    result.append([args[k][i] if i < len(args[k]) else missing_val for k in range(len(args))])
  return outList

## Sort a List of Dictionaries

dicts_lists = [
  {
    "Name": "James",
    "Age": 20,
  },
  {
     "Name": "May",
     "Age": 14,
  },
  {
    "Name": "Katy",
    "Age": 23,
  }
]

#There are different ways to sort that list
#1- Using the sort/ sorted function based on the age
dicts_lists.sort(key=lambda item: item.get("Age"))

#2- Using itemgetter module based on name
from operator import itemgetter
f = itemgetter('Name')
dicts_lists.sort(key=f)

## Sort a List of Strings

my_list = ["blue", "red", "green"]

#1- Using sort or srted directly or with specifc keys
my_list.sort() #sorts alphabetically or in an ascending order for numeric data 
my_list = sorted(my_list, key=len) #sorts the list based on the length of the strings from shortest to longest. 
# You can use reverse=True to flip the order

#2- Using locale and functools 
import locale
from functools import cmp_to_key
my_list = sorted(my_list, key=cmp_to_key(locale.strcoll)) 

## Sort a List Based on Another List

a = ['blue', 'green', 'orange', 'purple', 'yellow']
b = [3, 2, 5, 4, 1]
#Use list comprehensions to sort these lists
sortedList =  [val for (_, val) in sorted(zip(b, a), key=lambda x: x[0])]


## Map a List Into a Dictionary

mylist = ['blue', 'orange', 'green']
#Map the list into a dict using the map, zip and dict functions
mapped_dict = dict(zip(itr, map(fn, itr)))

## Merging Two or More Dictionaries

from collections import defaultdict
#merge two or more dicts using the collections module
def merge_dicts(*dicts):
  mdict = defaultdict(list)
  for d in dicts:
    for key in d:
      mdict[key].append(d[key])
  return dict(mdict)
  
## Inverting a Dictionary

my_dict = {
  "brand": "Ford",
  "model": "Mustang",
  "year": 1964
}
#Invert the dictionary based on its content
#1- If we know all values are unique.
my_inverted_dict = dict(map(reversed, my_dict.items()))

#2- If non-unique values exist
from collections import defaultdict
my_inverted_dict = defaultdict(list)
{my_inverted_dict[v].append(k) for k, v in my_dict.items()}

#3- If any of the values are not hashable
my_dict = {value: key for key in my_inverted_dict for value in my_inverted_dict[key]}

## Using F Strings

#Formatting strings with f string.
str_val = 'books'
num_val = 15
print(f'{num_val} {str_val}') # 15 books
print(f'{num_val % 2 = }') # 1
print(f'{str_val!r}') # books

#Dealing with floats
price_val = 5.18362
print(f'{price_val:.2f}') # 5.18

#Formatting dates
from datetime import datetime;
date_val = datetime.utcnow()
print(f'{date_val=:%Y-%m-%d}') # date_val=2021-09-24

## Checking for Substrings

addresses = ["123 Elm Street", "531 Oak Street", "678 Maple Street"]
street = "Elm Street"

#The top 2 methods to check if street in any of the items in the addresses list
#1- Using the find method
for address in addresses:
    if address.find(street) >= 0:
        print(address)
        
#2- Using the "in" keyword 
for address in addresses:
    if street in address:
        print(address)


## Get a String’s Size in Bytes

str1 = "hello"
str2 = "?"

def str_size(s):
  return len(s.encode('utf-8'))

str_size(str1)
str_size(str2)

## Checking if a File Exists

#Checking if a file exists in two ways
#1- Using the OS module
import os 
exists = os.path.isfile('/path/to/file')

#2- Use the pathlib module for a better performance
from pathlib import Path
config = Path('/path/to/file') 
if config.is_file(): 
    pass


# Parsing a Spreadsheet

import csv
csv_mapping_list = []
with open("/path/to/data.csv") as my_data:
    csv_reader = csv.reader(my_data, delimiter=",")
    line_count = 0
    for line in csv_reader:
        if line_count == 0:
            header = line
        else:
            row_dict = {key: value for key, value in zip(header, line)}
            csv_mapping_list.append(row_dict)
        line_count += 1

