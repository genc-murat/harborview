import { useState, useEffect } from 'react';
import PropTypes from 'prop-types';
import { Card, CardContent, CardHeader, CardTitle } from '@/components/ui/card';

const ContainerLogs = ({ containerId }) => {
  const [logs, setLogs] = useState('');

  useEffect(() => {
    const fetchLogs = async () => {
      try {
        const response = await fetch(`/containers/${containerId}/logs`);
        const data = await response.json();
        setLogs(data.logs);
      } catch {
        console.error('Failed to fetch logs');
      }
    };

    if (containerId) {
      fetchLogs();
    }
  }, [containerId]);

  return (
    <Card>
      <CardHeader>
        <CardTitle>Container Logs</CardTitle>
      </CardHeader>
      <CardContent>
        <pre className="bg-gray-900 text-gray-100 p-4 rounded-lg overflow-auto max-h-96">
          {logs || 'No logs available'}
        </pre>
      </CardContent>
    </Card>
  );
};

ContainerLogs.propTypes = {
  containerId: PropTypes.string.isRequired
};

export default ContainerLogs;