(in-ns 'wedge2.core)

(def sample-account {
	:name
	"bank/my_generic_bank"
	})

(def status-default
	"Default transaction status. If the date is in the past, the transaction is understood as status-done. If not, it is understood as status-planned."
	0)
(def status-done
	"Indicates that the transaction has already happened."
	1)
(def status-planned
	"Indicates that the transaction is still to happened."
	2)
(def status-cancelled
	"Indicates that the transaction was cancelled."
	-1)

(def sample-movement {
	:id	#uuid "00000000-0000-0000-0000-000000000000"
	:account "bank/my_generic_bank"
	:value (rationalize -10.27)
	:status	status-default
	:date (t/local-date 2013 3 20)
	})

(def sample-purchase {
	:id #uuid "00000000-0000-0000-0000-000000000000"
	:name "this brand of cheese"
	:generic-name "cheese/muzzarella"
	:unit-name "kg"; can be empty
	:qty (rationalize 0.2)
	:unit-cost (rationalize 50.28)
	:total-cost (rationalize 10.056)
	:period-start (t/local-date 2013 3 20)
	:period-end (t/local-date 2013 4 20)
	})

(def sample-transaction {
	:id #uuid "00000000-0000-0000-0000-000000000000"
	:name "steam stuff"
	:movements [sample-movement]
	:purchases [sample-movement]
	})

(defn movements [transactions]
	"Given a sequence of transactions, returns only the movements."
	(reduce concat (map :movements transactions)))

(defn non-cancelled-movements [transactions]
	"Given a sequence of transactions, returns only the non cancelled movements."
	(filter #(not= status-cancelled (:status %)) (movements transactions)))

(defn before [date movements]
	"Given a sequence of movements, returns only the ones that happened before the given date."
	(filter #(t/before? (:date %) date) movements))

(defn balance-on
	"Calculates the balance of an account at a given local date considering all movements given until the given date."
	[account-name movements date]
	(reduce +
		(map :value (t/before? date movements))))

(defn sort-movements [movements]
	"Given a list of movements, returns a new list sorting them from the oldest to the newest."
	(sort #(t/before? (:date %1) (:date %2)) movements))

(defn balances
	"Given a non-sorted list of movements, returns a map where each key is an account name and the values are vectors of maps ({:val ... :date ...}) indicating the balance of that account at that time."
	[movements]
	(loop [sorted (sort-movements movements) ret {}]
		(if (empty? sorted)
			ret
			(let [
				movement (first sorted)
				mov-date (:date movement)
				mov-val (:value movement 0)
				account (:account movement)
				history (get ret account [])
				last-val (:value (last history) 0)]
				(recur
					(rest sorted)
					(assoc ret account (conj history {:date mov-date :value (+ last-val mov-val)})))))))