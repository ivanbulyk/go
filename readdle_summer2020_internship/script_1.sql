/*
Find all current managers of each department and display his/her title, first name, last name, current salary.
*/
SELECT 
	   t.title,
	   e.first_name, e.last_name,
       s.salary

FROM employees.dept_manager dm

JOIN employees.titles t
     ON dm.emp_no = t.emp_no 
      AND dm.from_date = t.from_date


JOIN employees.employees e
     ON dm.emp_no = e.emp_no


JOIN employees.salaries s
     ON dm.emp_no = s.emp_no
      AND dm.to_date = s.to_date

WHERE dm.to_date = '9999-01-01'