/*
Find all departments, their current employee count, their current sum salary.
*/
SELECT 
      d.dept_name, de.dept_no,
      
      COUNT(e.emp_no) AS count_emp_no,
      SUM(s.salary) AS SUM_salary

FROM employees.employees e

JOIN employees.dept_emp de
     ON e.emp_no = de.emp_no
       AND e.hire_date = de.from_date
    
JOIN employees.salaries s
     ON de.emp_no = s.emp_no
      AND de.from_date = s.from_date

JOIN employees.departments d
     ON de.dept_no = d.dept_no

    WHERE de.to_date = '9999-01-01'
    GROUP BY de.dept_no