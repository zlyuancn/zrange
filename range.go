/*
-------------------------------------------------
   Author :       Zhang Fan
   date：         2019/5/9
   Description :
-------------------------------------------------
*/

package zrange

import (
    "container/list"
)

func Range(end int) chan int {
    return RangeStart(0, end, 1)
}

func RangeStart(start int, end int, step int) chan int {
    ch := make(chan int)
    go func() {
        for i := start; i != end; i += step {
            ch <- i
        }
        close(ch)
    }()
    return ch
}

func RangeList(l *list.List) (chan interface{}) {
    ch := make(chan interface{})
    go func() {
        e := l.Front()
        for {
            if e == nil {
                close(ch)
                return
            }
            ch <- e.Value
            e = e.Next()
        }
    }()
    return ch
}

func RangeListReverse(l *list.List) (chan interface{}) {
    ch := make(chan interface{})
    go func() {
        e := l.Back()
        for {
            if e == nil {
                close(ch)
                return
            }
            ch <- e.Value
            e = e.Prev()
        }
    }()
    return ch
}

func FRange(end int, fn func(i int)) {
    FRangeStart(0, end, 1, fn)
}

func FRangeStart(start int, end int, step int, fn func(i int)) {
    for i := start; i != end; i += step {
        fn(i)
    }
}

func FRangeList(l *list.List, fn func(i int, v interface{}) (stop bool)) {
    e := l.Front()
    var index int
    for {
        if e == nil || fn(index, e.Value) {
            return
        }

        index++
        e = e.Next()
    }
}

func FRangeListReverse(l *list.List, fn func(i int, v interface{}) (stop bool)) {
    e := l.Back()
    var index int
    for {
        if e == nil || fn(index, e.Value) {
            return
        }

        index++
        e = e.Prev()
    }
}
