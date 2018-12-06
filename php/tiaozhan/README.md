# PHP Coding Challenge Instructions

Write some code that will solve the following:

We want to add up the total number of seconds that a given object (represented
by the `SomeObject` class) was in a `RUNNING` state.  It doesn't matter for
this challenge what the object represents or what 'RUNNING' means.  A log is
kept of the changes to an object's state, and is stored in the `statusLog`
property of the object. The `statusLog` would be structured as follows:

```
array(
   array(
      'oldState' => STATE,
      'newState' => STATE,
      'date' => UNIX_TIMESTAMP
   ),
   array(
      'oldState' => STATE,
      'newState' => STATE,
      'date' => UNIX_TIMESTAMP
   )
);
```

There are three possible states, `PAUSED`, `RUNNING`, and `COMPLETE`.  The
array records the unix timestamp (seconds since epoch) of when the object went
into the `newState`, leaving `oldState`.  So we want a total sum of the time it
spent in a RUNNING state.   There may be several state changes, or thousands.
They are not necessarily in order.  An object may have never entered into a
RUNNING state.  No guarantees that once an object enters into a COMPLETE state
, that it won't then go back into RUNNING!  It's the wild west!

To further complicate things, the object may have hard defined `startTime` and
`stopTime` dates (also in seconds since epoch).  We may have state changes
outside of these dates, but want to make sure we only include run times within
these ranges (if given).  The start/stop dates are optional.  An object might
have a hard start, but no stop, or a hard stop, but no start.  Or none.  Or
both.  A null will be passed in to indicate the absence of a date.

In the `StateUtils` class, there is a `calculateTimeInState()` function where
you can begin implementing your solution. The function should take the array of
states, an optional start, and an optional stop as parameters, returning the
number of seconds within RUNNING state.

The code should not require a framework or any installations, so please do not
use any.

Within the attached php file you will find 12 test cases.  Each test case
includes an array of state changes, start dates and stop dates as well as
the expected results from the function.

Bonus challenge:  Make your function easily adapt to performing the same
calculation for the ‘PAUSED’ or ’COMPLETE’ states.
