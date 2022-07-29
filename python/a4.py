import csv

data_dict = dict()

with open('a.csv') as csvfile:
    csv_reader = csv.reader(csvfile)
    next(csv_reader) # ignore the title line
    for row in csv_reader:
        date = row[0]
        if date not in data_dict:
            data_dict[date] = []
        data_dict[date].append(tuple(row[1:]))

print(data_dict)

'''
Reported date,Gender,Age group,County
2022-11-07,Male,88+,City of London
2022-11-07,Female,30-39,City of London
2022-11-08,Male,70-79,City of London
2022-11-08,Male,98+,City of London
2022-11-09,Female,40-49,City of London
2022-11-09,Male,80-89,City of London

{
    '2022-11-07': [
        ('Male', '88+', 'City of London'), 
        ('Female', '30-39', 'City of London')
    ], 
    '2022-11-08': [
        ('Male', '70-79', 'City of London'), 
        ('Male', '98+', 'City of London')
    ], 
    '2022-11-09': [
        ('Female', '40-49', 'City of London'), 
        ('Male', '80-89', 'City of London')
    ]
}
'''
