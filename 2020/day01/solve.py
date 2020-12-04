with open('input_data.txt') as f:
  data = f.readlines()
data = [int(_.strip()) for _ in data]
data.sort()

for n1 in data:
  for n2 in data:
    for n3 in data:
      if n1 + n2 + n3 == 2020:
        print(n1, n2, n3, n2*n1*n3)
