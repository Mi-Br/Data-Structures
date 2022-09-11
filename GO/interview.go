// In a server, I would like to get running counters of how often an event occurred in the last second, the last minute, and the last hour.
// This could be used for a debug page on how much traffic there is right now:

// Page          | s   | m   | h
// ------------- | --- | --- | ----
// /index.html   | 10  | 400 | 5000
// /profile.html | 1   | 7   | 99

// This means that /index.html was viewed 10 times in the last second, 400 times in the last minute (i.e. the previous 60s) and 5000 times in the last hour (i.e. the previous 60min).
// The numbers should reflect what is happening on the server right now and be updated continuously.


type SecondMinuteHourCounter struct {
  secondsBuffer []int
  secondsCount int
  
}

func NewCounter() *SecondMinuteHourCounter {
  // TODO: implement
  var secondsBuffer []int = make([]int{}) // 3600 
  return &SecondMinuteHourCounter{secondsBuffer: secondsBuffer}   
}

// advance is called once per second by an external clock system.
func (c *SecondMinuteHourCounter) advance() {
  // TODO: implement
  c.secondsCount ++
  c.secondsBuffer = append(c.secondsBuffer, 0)
  l:=len(c)
  if l == 3601 {
    c.secondsBuffer = c.secondsBuffer[1:]
  }
}

// inc is called to record a new event.
func (c *SecondMinuteHourCounter) inc() {
  // TODO: implement
  l:=len(c)
  if  l == 0 {
    c.secondsBuffer = append(c.secondsBuffer, 1) 
  } else {
    c.secondsBuffer[l-1] = c.secondsBuffer[l-1]+1
  }
}

// getCounter retrieves the number of events occured in the last second, minute or hour.
// TODO: update the function signature as needed.
func (c *SecondMinuteHourCounter) getCounter() int {
  // TODO: implement
}


c := NewCounter()
c.inc()
c.inc()
c.advance()
c.inc()
c.getCounter('second') --> 1
c.getCounter('minute') --> 3


Events per second: 1, 1, 1, 1 ... 1 (60 times) 
g.getCounter('minute') --> 60
g.getCounter('second') --> 1
Events per second: 1, 1, 1, 1 ... 1 (60 times), 2
g.getCounter('minute') --> 61
g.getCounter('second') --> 2
Events per second: 1, '1', 1, 1 ... 1 (60 times), 2, 5
g.getCounter('minute') --> 61 - 1 + 5 = 65


// Pseudecode

c := NewCounter()
EverySecond: c.advance()

func UserLoadEvent() {
  c.inc()
}

func ShowStatistics() {
  s := c.getCounter('second')
  fmt.Printf("User events in last second: %v", s)
}

