Create extension if not exists "uuid-ossp";

CREATE TABLE departments (
    id uuid PRIMARY KEY DEFAULT uuid_generate_v4(),
    name VARCHAR(100)
);

CREATE TABLE employees (
    id uuid PRIMARY KEY DEFAULT uuid_generate_v4(),
    name VARCHAR(100),
    department_id uuid REFERENCES departments(id)
);

CREATE TABLE salaries (
    id uuid PRIMARY KEY DEFAULT uuid_generate_v4(),
    employee_id uuid REFERENCES employees(id),
    amount DECIMAL(10, 2),
    date DATE
);

INSERT INTO departments (name) VALUES
    ('Engineering'),
    ('Marketing'),
    ('Finance'),
    ('Human Resources'),
    ('Sales'),
    ('Accounting'),
    ('Legal');

INSERT INTO employees (name, department_id) VALUES
    ('Bob Stoke', (select id from departments where name = 'Engineering')),
    ('Alice Johnson', (select id from departments where name = 'Marketing')),
    ('Jane Carlen', (select id from departments where name = 'Finance')),
    ('Mike Jones', (select id from departments where name = 'Human Resources')),
    ('Mary Brown', (select id from departments where name = 'Sales')),
    ('Bob Williams', (select id from departments where name = 'Accounting')),
    ('Eva Davis', (select id from departments where name = 'Legal'));

INSERT INTO salaries (employee_id, amount, date) VALUES
    ((select id from employees where name = 'Bob Stoke'), 80000.00, '2023-11-30'),
    ((select id from employees where name = 'Alice Johnson'), 75000.00, '2023-11-30'),
    ((select id from employees where name = 'Jane Carlen'), 90000.00, '2023-10-30'),
    ((select id from employees where name = 'Mike Jones'), 82000.00, '2024-12-30'),
    ((select id from employees where name = 'Mary Brown'), 85000.00, '2023-11-30'),
    ((select id from employees where name = 'Bob Williams'), 95000.00, '2023-10-30'),
    ((select id from employees where name = 'Eva Davis'), 100000.00, '2023-12-30'),
    ((select id from employees where name = 'Jane Carlen'), 90000.00, '2023-11-30'),
    ((select id from employees where name = 'Alice Johnson'), 75000.00, '2023-12-30'),
    ((select id from employees where name = 'Mary Brown'), 85000.00, '2023-12-30'),
    ((select id from employees where name = 'Bob Stoke'), 80000.00, '2023-12-30'),
    ((select id from employees where name = 'Jane Carlen'), 90000.00, '2023-12-30'),
    ((select id from employees where name = 'Bob Williams'), 95000.00, '2023-11-30'),
    ((select id from employees where name = 'Mike Jones'), 82000.00, '2024-01-23');
