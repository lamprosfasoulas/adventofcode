# Open the file and read lines
with open("tinyinput.txt") as f:
    lines = [list(line.strip()) for line in f]  # Split each line into characters

# Transpose the content
transposed = zip(*lines)

# Write the transposed output to a new file
with open("transposed.txt", "w") as out:
    for row in transposed:
        out.write("".join(row) + "\n")

