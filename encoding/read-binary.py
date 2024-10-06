filename = "students.bin"


file = open(filename, "rb")

content = file.read()

result = []

for i in range(0, len(content), 6):

    studentId = content[i]
    maths = content[i + 1]
    science = content[i +2]
    social = content[i + 3]
    english = content[i + 4]
    tamil = content[i + 5]
    print(f"StudentId: {studentId}, Maths: {maths}, Science: {science}, Social: {social}, English: {english}, Tamil: {tamil}")


