filename = "students.csv"


file = open(filename, "r")

content = file.read()
lines = content.splitlines()

result = []

first = True
for line in lines:
    if first:
        first = False
        continue
    values = line.split(",")
    print(f"StudentId: {values[0]}, Maths: {values[1]}, Science: {values[2]}, Social: {values[3]}, English: {values[4]}, Tamil: {values[5]}")







