import matplotlib.pyplot as plt

x = []
list = []

with open("Azul.txt") as archivo:
    for line in archivo:
      list.append(int(line))
      
for i in range (256):
  x.append(i+1)

plt.hist(list, bins=255, color = "blue", rwidth=1)
plt.title("Histograma de intensidad BLUE")
plt.xlabel("Intensidad de color")
plt.ylabel("Cant Pixeles")
plt.show()