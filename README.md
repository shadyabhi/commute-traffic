# commute-traffic

It's a WIP.

## Config file

	apiKeys:
		# Get API key from https://developers.google.com/maps/documentation/directions/get-api-key
		- secret-api-key

	tracks:
		- source: Zonasha Elegance, Harlur Main Road
		  destination: LinkedIn, Global Technology Park
		- source: Zonasha Elegance, Harlur Main Road
		  destination: Slurp Cafe, Indiranagar

## Run

Once it's installed, you can run this as a periodic job. A sample run looks like:-

	➜ $?=0 @arastogi-mn3 golang/commute-traffic [11:47AM] (master↑1|✚1…)➤ go run *.go
	2017/11/12 11:47:39 Currently working on, SRC=Zonasha Elegance, Harlur Main Road, DST=LinkedIn, Global Technology Park
	2017/11/12 11:47:41 Currently working on, SRC=Zonasha Elegance, Harlur Main Road, DST=Slurp Cafe, Indiranagar
	2017/11/12 11:47:41 Indexed to ES: doc={%!s(int64=6787) 1510467459 %!s(int64=1213) Reliable Silver Oak Layout, Eastwood Township, Harlur, Bengaluru, Karnataka 560035, India 7th floor,Tower A, Global Technology Park, Devarabisanahalli, Next to Intel, Adarsh Palm Retreat, Bengaluru, Karnataka 560103, India %!s(int64=1168)}
	2017/11/12 11:47:41 Indexed to ES: doc={%!s(int64=14363) 1510467461 %!s(int64=2677) Reliable Silver Oak Layout, Eastwood Township, Harlur, Bengaluru, Karnataka 560035, India 1079, 12th Main Road, HAL 2nd Stage, Indiranagar, Bengaluru, Karnataka 560038, India %!s(int64=2584)}

	>>>  3s elasped...
	➜ $?=0 @arastogi-mn3 golang/commute-traffic [11:47AM] (master↑1|✚1…)➤
