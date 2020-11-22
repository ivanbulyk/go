/*
Find all employees (department, title, first name, last name, hire date, how many years they have been working) 
to congratulate them on their hire anniversary this month.
*/
SELECT 
      d.dept_name,
      t.title,
      e.first_name,
      e.last_name,
      e.hire_date,
      TIMESTAMPDIFF(YEAR,e.hire_date,CURDATE())  AS Experience 
FROM employees.employees e
JOIN employees.titles t
     ON e.emp_no = t.emp_no
      AND e.hire_date = t.from_date
JOIN employees.dept_emp de
     ON e.emp_no = de.emp_no
JOIN employees.departments d
     ON de.dept_no = d.dept_no
   
 