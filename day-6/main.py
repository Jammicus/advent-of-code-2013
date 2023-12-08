def main():

    combined = parseFile()
    partOne(combined)
    partTwo(combined)

def parseFile():
    times = []
    distance = []
    combined = []

    with open('day6.txt') as f:
        firstLine = f.readline().rstrip('\n')
        splitLine = firstLine.split(" ")
        for entry in splitLine:
            if entry.isnumeric():
                times.append(entry)
        
        secondLine = f.readline().rstrip('\n')
        secondSplitLine = secondLine.split(" ")
        for entry in secondSplitLine:
            if entry.isnumeric():
                distance.append(entry)

    
        for x in range(len(times)):
            combined.append((times[x], distance[x]))
        
    f.close()

    return combined




def partOne(times):
    toMultiply = []
    for entry in times:
        time = int(entry[0])
        distance = int(entry[1])
        counter = 0

        for run in range(time):

            millisecondsHeld = run 
            if millisecondsHeld == 0:
                continue
            movingTime = time - millisecondsHeld
            result = movingTime * millisecondsHeld

            if result > distance:
                counter = counter +1
        
        toMultiply.append(counter)
    
    res = 1
    for entry in toMultiply:
        res = res * entry
    
    print(res)


def partTwo(times):
    toMultiply = []
    time = ""
    distance = ""
    counter = 0
    for entry in times:
        time = f"{time}{entry[0]}"
        distance = f"{distance}{entry[1]}"

    time = int(time)
    distance = int(distance)

    for run in range(time):

        millisecondsHeld = run 
        if millisecondsHeld == 0:
            continue
        # 6 =  7 -1
        movingTime = time - millisecondsHeld
        result = movingTime * millisecondsHeld

        if result > distance:
            counter = counter +1
    
    toMultiply.append(counter)

    res = 1
    for entry in toMultiply:
        res = res * entry

    print(res)

if __name__ == "__main__":
    main()