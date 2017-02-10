package quandl

var dbSetup bool

func Set() {
	/*influxdb.Setup(config.Database{
		User:     "trades",
		Password: "xEoQGTgTl7UEqGESSPaL",
		Host:     "http://192.168.80.131:8086",
		Schema:   "markets",
	})*/
	//dbSetup = true
}

/*func TestShouldGetNew001(t *testing.T) {
	if !dbSetup {
		Set()
	}
	now := time.Now().UTC()

	c := code{
		frequency:   "daily",
		lastUpdated: now,
	}
	now.Day()
	c.ShouldGetNew()
	yesterday := now.Add(-day)
	yesterday.Day()
	d := code{
		frequency:   "daily",
		lastUpdated: yesterday,
	}

	d.ShouldGetNew()

}

func TestShouldGetNew002(t *testing.T) {
	if !dbSetup {
		Set()
	}
	now := time.Now().UTC()

	c := code{
		frequency:   "weekly",
		lastUpdated: now,
	}
	now.Day()
	c.ShouldGetNew()
	yesterday := now.Add(-day * 7)
	yesterday.Day()
	d := code{
		frequency:   "weekly",
		lastUpdated: yesterday,
	}

	d.ShouldGetNew()

}

func TestShouldGetNew003(t *testing.T) {
	if !dbSetup {
		Set()
	}
	c := code{
		frequency: "weekly",
	}
	c.ShouldGetNew()

	d := code{
		lastUpdated: time.Now(),
	}

	d.ShouldGetNew()

}

func TestShouldGetNew004(t *testing.T) {
	if !dbSetup {
		Set()
	}
	now := time.Now().UTC()

	c := code{
		frequency:   "monthly",
		lastUpdated: now,
	}
	c.ShouldGetNew()
	yesterday := now.Add(-day * 32)
	d := code{
		frequency:   "monthly",
		lastUpdated: yesterday,
	}

	d.ShouldGetNew()

}

func TestShouldGetNew005(t *testing.T) {
	if !dbSetup {
		Set()
	}
	now := time.Now().UTC()

	c := code{
		frequency:   "quarterly",
		lastUpdated: now,
	}
	c.ShouldGetNew()
	yesterday := now.Add(-day * 100)
	d := code{
		frequency:   "quarterly",
		lastUpdated: yesterday,
	}

	d.ShouldGetNew()

}

func TestShouldGetNew006(t *testing.T) {
	if !dbSetup {
		Set()
	}
	now := time.Now().UTC()

	c := code{
		frequency:   "yearly",
		lastUpdated: now,
	}
	now.Day()
	c.ShouldGetNew()
	yesterday := now.Add(-day * 366)
	yesterday.Day()
	d := code{
		frequency:   "yearly",
		lastUpdated: yesterday,
	}

	d.ShouldGetNew()

}

func TestShouldGetNew007(t *testing.T) {
	if !dbSetup {
		Set()
	}
	now := time.Now().UTC()

	c := code{
		frequency:   "yearly",
		lastUpdated: now,
	}

	c.Endpoint()
	c.Frequency()
	c.Update(now.Add(day))
	c.SetFrequency("daily")

}

func TestShouldGetNew008(t *testing.T) {
	if !dbSetup {
		Set()
	}
	now := time.Now().UTC()

	c := code{
		frequency:   "annually",
		lastUpdated: now,
	}
	now.Day()
	c.ShouldGetNew()
	yesterday := now.Add(-day * 366)
	yesterday.Day()
	d := code{
		frequency:   "annually",
		lastUpdated: yesterday,
	}

	d.ShouldGetNew()

}

func TestShouldGetNew009(t *testing.T) {
	if !dbSetup {
		Set()
	}
	now := time.Now().UTC()

	c := code{
		frequency:   "asdf",
		lastUpdated: now,
	}
	now.Day()
	c.ShouldGetNew()
	yesterday := now.Add(-day * 366)
	yesterday.Day()
	d := code{
		frequency:   "asdf",
		lastUpdated: yesterday,
	}

	d.ShouldGetNew()

}*/
