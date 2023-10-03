-- Create the database if it doesn't exist
CREATE DATABASE IF NOT EXISTS awsDB;

-- Use the awsDB database
USE awsDB;

-- Create a sample table
CREATE TABLE IF NOT EXISTS ec2_instances (
                                             id INT AUTO_INCREMENT PRIMARY KEY,
                                             instance_id VARCHAR(255) NOT NULL
    );

-- Insert some initial data
INSERT INTO ec2_instances (instance_id) VALUES
                                            ('i-1234567890abcdef0'),
                                            ('i-9876543210abcdef1');
