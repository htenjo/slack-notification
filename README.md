# slack-notification

This tool allows you to retrieve relevant insights from a 
given repository in Github and send it to a slack channel
previously configured.

Metrics we can send:
- Open PRs
    - Red: Created time > 6h
    - Black: Created time 4h - 6h
    - Green: Created time < 4h
