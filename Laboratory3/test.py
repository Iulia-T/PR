import requests

BASE = "http://127.0.0.1:5000/"

data = [{"name": "pizza", "preparation-time": 20 , "complexity": 2, "cooking-apparatus": "oven"},
        {"name": "salad", "preparation-time": 10 , "complexity": 1, "cooking-apparatus": "none"},
        {"name": "zeama", "preparation-time": 7 , "complexity": 1, "cooking-apparatus": "stove"},
        {"name": "Scallop Sashimi with Meyer Lemon Confit", "preparation-time": 32 , "complexity": 3, "cooking-apparatus": "none"},
        {"name": "Island Duck with Mulberry Mustard", "preparation-time": 25 , "complexity": 3, "cooking-apparatus": "oven"},
        {"name": "Waffles", "preparation-time": 10 , "complexity": 1, "cooking-apparatus": "stove"}]


for i in range(len(data)):
    response = requests.put(BASE + "food/" + str(i), data[i])
    print(response.json())

input()

response = requests.delete(BASE + "food/1")
print(response)

input()

for i in range(len(data)):
    response = requests.get(BASE + "food/" + str(i))
    print(response.json())
    
input()

for i in range(len(data)):
    for sub in data[i]:
        if sub['Name' == 'zeama']:
            response = requests.put(BASE + "food/" + str(i), {"name": "zeama", "preparation-time": 5 , "complexity": 2, "cooking-apparatus": "stove"})
            break
print(response.json())
    
        