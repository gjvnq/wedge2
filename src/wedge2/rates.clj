(in-ns 'wedge2.core)

; Alphanumeric only assets (BTC, BRL, USD, DASH, ...) are interpreted as currencies to be converted via coinmarketcap.com
; If there is a semi-colon (:), the first part is understood as te name of an exchange and second is understood as a stock name. Ex: BOVESPA:PETR4, NYSE:GOOG
; Assets beginning in a hyphen (!) are understood as 'manual assets', whose price must be set manually 

(load "rates-coinmarketcap")
(defn conv-bovespa
	[{date :date asset-src :asset-src}]
	{:error "not implemented" :factor 0 :asset "BRL"})
(defn conv-nyse
	[{date :date asset-src :asset-src}]
	{:error "not implemented" :factor 0 :asset "USD"})
(defn conv-err
	[m]
	{:error "not implemented" :factor 0 :asset "USD"})
(defn conv-manual
	[{date :date asset-src :asset-src historic :manual-historic}]
	{:error "not implemented" :factor 0 :asset ""})

(def asset-converters-regex [
	{:regex #"([A-Z0-9a-z]+)" :func conv-coinmarketcap}
	{:regex #"B3:([A-Z0-9a-z]+)" :func conv-bovespa}; B3 (former BM&FBOVESPA)
	{:regex #"NYSE:([A-Z0-9a-z]+)" :func conv-nyse}; New York Stock Exchange
	{:regex #"NASDAQ:([A-Z0-9a-z]+)" :func conv-err}; Nasdaq Stock Market
	{:regex #"LSEG:([A-Z0-9a-z]+)" :func conv-err}; London Stock Exchange Group
	{:regex #"JPX:([A-Z0-9a-z]+)" :func conv-err}; Japan Exchange
	{:regex #"EURONEXT:([A-Z0-9a-z]+)" :func conv-err}; Euronext NV
	{:regex #"!([A-Z0-9a-z]+)" :func conv-manual}])

(defn asset-converter
	"Given an asset, returns the function used to convert it or nil if not found."
	([asset]
	(asset-converter asset asset-converters-regex))
	
	([asset the-regex-list]
	(loop [regex-list the-regex-list]
		(let [
			head (first regex-list) 
			this-regex (:regex head)
			this-func (:func head)]
			(if (nil? head)
				nil
				(if (re-matches this-regex asset)
					this-func
					(recur (rest regex-list))))))))

(defn non-cached-asset-factor
	([asset-src asset-dst date]
	(let [ans (non-cached-asset-factor asset-src asset-dst date #{})]
		(if ans
			ans
			; If we ccant convert one way, try the inverse way
			(/ 1 (:factor (non-cached-asset-factor asset-dst asset-src date #{}))))))

	([asset-src asset-dst date tried]
	(if (= asset-src asset-dst)
		1; "Obviously" - A High Functioning Sociopath
		(let [
			ans ((asset-converter asset-src) {:asset-src asset-src :asset-dst asset-dst :date date}); Try to convert asset-src into asset-dst
			asset-ans (:asset ans)]; Stocks may return in an asset different froma asset-dst
			(if (= asset-ans asset-dst)
				(:factor ans); We got the right answer, job done
				(if (contains? tried [asset-src asset-ans])
					nil; Do not try the same path twice
					(non-cached-asset-factor asset-ans asset-dst))))))); Let us try a new path
