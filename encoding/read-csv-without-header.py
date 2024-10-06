filename = "students-without-header.csv"


file = open(filename, "r")

content = file.read()
lines = content.splitlines()

result = []

binary_file = open("students.bin", "wb")

for line in lines:
    values = line.split(",")
    print(f"StudentId: {values[0]}, Maths: {values[1]}, Science: {values[2]}, Social: {values[3]}, English: {values[4]}, Tamil: {values[5]}")
    values_as_bytes = [int(values[0]), int(values[1]), int(values[2]), int(values[3]), int(values[4]), int(values[5])]
    binary_file.write(bytes(values_as_bytes))



file.close()

