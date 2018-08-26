type TimePoint struct {
    Time int
    Sub int
    Next *TimePoint
}

func findMinDifference(timePoints []string) int {
    var root, preNode, node *TimePoint
    var hour, minute, time int
    var min = 24*60
    for _, point := range timePoints {
        tt := strings.Split(point, ":")
        hour, _ = strconv.Atoi(tt[0])
        minute, _ = strconv.Atoi(tt[1])
        time = 60*hour + minute
        if root == nil {
            root = new(TimePoint)
            root.Time = time
            root.Sub = 24*60
            continue
        }
        preNode = nil
        for node = root; node.Next != nil && node.Time < time ; node = node.Next {
            preNode = node
        }
        if node.Time == time {
            return 0
        }
        if (preNode == nil && node.Next != nil) || (preNode == nil && node.Next == nil && node.Time > time) {
            preNode = new(TimePoint)
            preNode.Time = time
            preNode.Sub = root.Time - time
            preNode.Next = root
            root = preNode
        } else if node.Next == nil && node.Time <= time {
            node.Next = new(TimePoint)
            node.Next.Time = time
            node.Sub = time - node.Time
        } else {
            preNode.Next = new(TimePoint)
            preNode.Sub = time - preNode.Time
            preNode.Next.Next = node
            preNode.Next.Time = time
            preNode.Next.Sub = node.Time - time
        }
    }
    for node = root; node.Next != nil; node = node.Next {
        if min > node.Sub {
            min = node.Sub
        }
    }
    if min > 24*60 - node.Time + root.Time {
        min =  24*60 - node.Time + root.Time
    }
    return min
}
