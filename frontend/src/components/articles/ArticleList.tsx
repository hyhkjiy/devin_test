import { useState, useEffect } from 'react';
import { Card, CardContent, CardHeader, CardTitle } from "../ui/card";
import { Button } from "../ui/button";
import { Alert, AlertDescription } from "../ui/alert";

interface Article {
  id: number;
  title: string;
  content: string;
  publishTime: string;
}

export function ArticleList() {
  const [articles, setArticles] = useState<Article[]>([]);
  const [error, setError] = useState('');
  const [selectedArticle, setSelectedArticle] = useState<Article | null>(null);

  useEffect(() => {
    fetchArticles();
  }, []);

  const fetchArticles = async () => {
    try {
      const response = await fetch('http://localhost:8000/api/articles');
      if (!response.ok) {
        throw new Error('Failed to fetch articles');
      }
      const data = await response.json();
      setArticles(data);
      setError('');
    } catch (err) {
      setError('Failed to load articles. Please try again later.');
    }
  };

  const formatDate = (dateString: string) => {
    return new Date(dateString).toLocaleDateString('en-US', {
      year: 'numeric',
      month: 'long',
      day: 'numeric',
      hour: '2-digit',
      minute: '2-digit'
    });
  };

  return (
    <div className="space-y-6">
      {error && (
        <Alert variant="destructive">
          <AlertDescription>{error}</AlertDescription>
        </Alert>
      )}
      
      {selectedArticle ? (
        <Card>
          <CardHeader>
            <CardTitle className="text-2xl font-bold">{selectedArticle.title}</CardTitle>
            <p className="text-sm text-gray-500">
              Published on {formatDate(selectedArticle.publishTime)}
            </p>
          </CardHeader>
          <CardContent>
            <p className="whitespace-pre-wrap">{selectedArticle.content}</p>
            <Button
              variant="outline"
              className="mt-4"
              onClick={() => setSelectedArticle(null)}
            >
              Back to List
            </Button>
          </CardContent>
        </Card>
      ) : (
        <div className="grid gap-4">
          {articles.map((article) => (
            <Card key={article.id} className="cursor-pointer hover:bg-gray-50" onClick={() => setSelectedArticle(article)}>
              <CardHeader>
                <CardTitle>{article.title}</CardTitle>
                <p className="text-sm text-gray-500">
                  Published on {formatDate(article.publishTime)}
                </p>
              </CardHeader>
            </Card>
          ))}
        </div>
      )}
    </div>
  );
}
