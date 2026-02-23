import csv
import matplotlib.pyplot as plt

workers = []
times = []

with open("scaling.csv") as f:
    reader = csv.DictReader(f)
    for row in reader:
        workers.append(int(row["workers"]))
        times.append(float(row["time"]))

plt.figure()
plt.plot(workers, times, marker="o")
plt.xlabel("Number of Workers")
plt.ylabel("Execution Time (seconds)")
plt.title("Concurrent Log Processor Scaling")
plt.grid(True)

plt.savefig("scaling.png")
plt.show()
