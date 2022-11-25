import matplotlib.pyplot as plt

x = []
list = []

with open("Verde.txt") as archivo:
    for line in archivo:
      list.append(int(line))
      
for i in range (256):
  x.append(i+1)

plt.hist(list, bins=255, color = "green", rwidth=1)
plt.title("Histograma de intensidad GREEN")
plt.xlabel("Intensidad de color")
plt.ylabel("Cant Pixeles")
plt.show()