---
name: Bug report
about: "Create a bug report to help us improve"
---

<!--
Please answer all the questions with enough information. All issues not following this template will be closed immediately.
If you are not sure if your question is truely a bug of Vdonkey, please discuss it at https://github.com/vdonkey/accelerator/discussions first.
-->

## What version of Vdonkey are you using?

<!-- If you deploy different versions of Vdonkey on server and client, please explicitly point out -->


## What's your scenario of using Vdonkey?

<!-- E.g., watching YouTube videos in browsers via Socks/VMess proxy -->


## What problems have you encountered?

<!-- Please describe in detail, such as timeout, fake TLS certificate, etc -->


## What's your expectation?



## Please attach your configuration here

<!-- Remember to mask your IP address or hostname -->

**Server configuration:**

```javascript
// Please attach your server configuration here.

```

**Client configuration:**

```javascript
// Please attach your client configuration here.

```

## Please attach error logs here

<!--
only trailing lines if the log file is large in size.
Error log file is usually at `/var/log/accelerator/error.log` on Linux.
-->

**Server error log:**

```javascript
// Please attach your server error log here.

```

**Client error log:**

```javascript
// Please attach your client error log here.

```

## Please attach access log here

<!-- Access log is usually at '/var/log/accelerator/access.log' on Linux. -->

```javascript
// Please attach your server access log here.

```

## Other configurations (such as Nginx) and logs here



## If Vdonkey cannot start up, please attach output from `--test` command



## If Vdonkey service is abnormal, please attach journal log here

<!-- Usual command is `journalctl -u accelerator` -->

<!-- Please review your issue and check the format before submitting. -->
