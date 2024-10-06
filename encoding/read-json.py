import json
filename = "students.json"


file = open(filename, "r")

content = file.read()
records = json.loads(content)

for record in records:
    studentId = record["StudentId"]
    maths = record["Maths"]
    tamil = record["Tamil"]
    social = record["Social"]
    science = record["Science"]
    english = record["English"]
    print(f"StudentId: {studentId}, Maths: {maths}, Science: {science}, Social: {social}, English: {english}, Tamil: {tamil}")
