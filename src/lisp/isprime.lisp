(defun primep (n)
  "Is N prime?"
  (and (> n 1)
       (or (= n 2) (oddp n))
       (loop for i from 3 to (isqrt n) by 2
	  never (zerop (rem n i)))))

(print (primep 7))

; easier, but slower implementation

(defun isPrime (n)
  (setq isPrime T)
  (loop for i from 3 to (- n 1) do
    (setq modISNull (zerop (rem n i)))
    (if modIsNull (setq isPrime nil))
  )
  isPrime
)

(print (isPrime 17))