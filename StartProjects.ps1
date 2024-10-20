# StartProjects.ps1
Start-Process powershell -ArgumentList "Start-Process powershell -ArgumentList 'cd ./backend; air; Pause'"
Start-Process powershell -ArgumentList "Start-Process powershell -ArgumentList 'cd ./backend/src/python; python visualize_performance.py; Pause'"
Start-Process powershell -ArgumentList "Start-Process powershell -ArgumentList 'cd ./frontend; npm start; Pause'"
