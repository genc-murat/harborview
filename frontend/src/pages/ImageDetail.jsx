import { useState, useEffect } from 'react';
import { useParams } from 'react-router-dom';
import { Card, CardContent, CardHeader, CardTitle } from '@/components/ui/card';

const ImageDetail = () => {
  const { id } = useParams();
  const [image, setImage] = useState(null);
  const [loading, setLoading] = useState(true);

  useEffect(() => {
    const fetchImageDetails = async () => {
      try {
        const response = await fetch(`/images/${id}/inspect`);
        const data = await response.json();
        setImage(data);
      } catch {
        console.error('Failed to fetch image details');
      } finally {
        setLoading(false);
      }
    };

    fetchImageDetails();
  }, [id]);

  if (loading) {
    return (
      <div className="flex justify-center items-center h-full">
        <div className="animate-spin rounded-full h-12 w-12 border-b-2 border-gray-900" />
      </div>
    );
  }

  if (!image) {
    return (
      <div className="text-center">
        <h2 className="text-xl">Image not found</h2>
      </div>
    );
  }

  return (
    <div className="space-y-6">
      <h1 className="text-2xl font-bold">Image Details</h1>
      
      <div className="grid gap-6 md:grid-cols-2">
        <Card>
          <CardHeader>
            <CardTitle>Basic Info</CardTitle>
          </CardHeader>
          <CardContent>
            <dl className="space-y-2">
              <div>
                <dt className="font-medium">ID</dt>
                <dd className="text-gray-600">{image.Id}</dd>
              </div>
              <div>
                <dt className="font-medium">Created</dt>
                <dd className="text-gray-600">
                  {new Date(image.Created).toLocaleString()}
                </dd>
              </div>
              <div>
                <dt className="font-medium">Size</dt>
                <dd className="text-gray-600">
                  {(image.Size / 1024 / 1024).toFixed(2)} MB
                </dd>
              </div>
            </dl>
          </CardContent>
        </Card>

        <Card>
          <CardHeader>
            <CardTitle>Configuration</CardTitle>
          </CardHeader>
          <CardContent>
            <dl className="space-y-2">
              <div>
                <dt className="font-medium">Architecture</dt>
                <dd className="text-gray-600">{image.Architecture}</dd>
              </div>
              <div>
                <dt className="font-medium">OS</dt>
                <dd className="text-gray-600">{image.Os}</dd>
              </div>
              <div>
                <dt className="font-medium">Docker Version</dt>
                <dd className="text-gray-600">{image.DockerVersion}</dd>
              </div>
            </dl>
          </CardContent>
        </Card>
      </div>
    </div>
  );
};

export default ImageDetail;