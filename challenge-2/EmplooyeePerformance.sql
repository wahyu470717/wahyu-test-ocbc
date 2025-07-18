SELECT 
    CONCAT(
        FullName, 
        ' (ID: ', 
        EmployeeID, 
        ') has a performance rating of: ',
        CASE 
            WHEN PerformanceScore < 50 THEN 'Needs Improvement'
            WHEN PerformanceScore >= 50 AND PerformanceScore < 75 THEN 'Meets Expectations'
            WHEN PerformanceScore >= 75 AND PerformanceScore < 90 THEN 'Exceeds Expectations'
            WHEN PerformanceScore >= 90 THEN 'Outstanding'
            ELSE 'Unknown'
        END
    ) AS PerformanceReport
FROM EMPLOYEES
ORDER BY EmployeeID ASC;
