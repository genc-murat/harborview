import { useState, useEffect } from 'react';
import { useParams } from 'react-router-dom';
import { Card, CardContent, CardHeader, CardTitle } from '@/components/ui/card';
import ContainerLogs from '@/components/containers/ContainerLogs';

const ContainerDetail = () => {
  const { id } = useParams();
  const [container, setContainer] = useState(null);
  const [loading, setLoading] = useState(true);

  useEffect(() => {
    const fetchContainerDetails = async () => {
      try {
        const response = await fetch(`/containers/${id}`);
        const data = await response.json();
        setContainer(data);
      } catch {
        console.error('Failed to fetch container details');
      } finally {
        setLoading(false);
      }
    };

    fetchContainerDetails();
  }, [id]);

  if (loading) {
    return (
      <div className="flex justify-center items-center h-full">
        <div className="animate-spin rounded-full h-12 w-12 border-b-2 border-gray-900" />
      </div>
    );
  }

  if (!container) {
    return (
      <div className="text-center">
        <h2 className="text-xl">Container not found</h2>
      </div>
    );
  }

  return (
    <div className="space-y-6">
      <h1 className="text-2xl font-bold">Container Details</h1>
      
      <div className="grid gap-6 md:grid-cols-2">
        <Card>
          <CardHeader>
            <CardTitle>Info</CardTitle>
          </CardHeader>
          <CardContent>
            <dl className="space-y-2">
              <div>
                <dt className="font-medium">ID</dt>
                <dd className="text-gray-600">{container.Id}</dd>
              </div>
              <div>
                <dt className="font-medium">Name</dt>
                <dd className="text-gray-600">{container.Name}</dd>
              </div>
              <div>
                <dt className="font-medium">Status</dt>
                <dd className="text-gray-600">{container.State?.Status}</dd>
              </div>
              <div>
                <dt className="font-medium">Image</dt>
                <dd className="text-gray-600">{container.Config?.Image}</dd>
              </div>
            </dl>
          </CardContent>
        </Card>

        <Card>
          <CardHeader>
            <CardTitle>Network</CardTitle>
          </CardHeader>
          <CardContent>
            <dl className="space-y-2">
              <div>
                <dt className="font-medium">IP Address</dt>
                <dd className="text-gray-600">
                  {container.NetworkSettings?.IPAddress || 'Not assigned'}
                </dd>
              </div>
              <div>
                <dt className="font-medium">Ports</dt>
                <dd className="text-gray-600">
                  {Object.keys(container.NetworkSettings?.Ports || {}).join(', ') || 'No ports exposed'}
                </dd>
              </div>
            </dl>
          </CardContent>
        </Card>
      </div>

      <ContainerLogs containerId={id} />
    </div>
  );
};

export default ContainerDetail;