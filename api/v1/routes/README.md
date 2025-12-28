import os
from setuptools import setup, find_packages

with open('requirements.txt', 'r') as f:
    requirements = f.read().splitlines()

setup(
    name='analytics-worker',
    version='1.0.0',
    description='A worker for analytics tasks',
    author='Your Name',
    author_email='your@email.com',
    packages=find_packages(),
    install_requires=requirements,
    entry_points={
        'console_scripts': ['analytics-worker=analytics_worker.main:main'],
    },
    classifiers=[
        'Development Status :: 5 - Production/Stable',
        'Intended Audience :: Developers',
        'License :: OSI Approved :: MIT License',
        'Programming Language :: Python :: 3',
        'Programming Language :: Python :: 3.6',
        'Programming Language :: Python :: 3.7',
        'Programming Language :: Python :: 3.8',
    ],
)