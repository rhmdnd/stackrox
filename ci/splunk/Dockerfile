FROM splunk/splunk:8.2.4
COPY server.conf /opt/splunk/etc/system/local/server.conf
COPY inputs.conf /opt/splunk/etc/apps/splunk_httpinput/local/inputs.conf

RUN sudo mkdir -p /opt/splunk/etc/apps/stackrox/default
COPY syslog_inputs.conf /opt/splunk/etc/apps/stackrox/default/inputs.conf
