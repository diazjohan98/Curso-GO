with open("datos.csv", "r") as f:
    lineas = sum(1 for _ in f)
print("Líneas:", lineas)
