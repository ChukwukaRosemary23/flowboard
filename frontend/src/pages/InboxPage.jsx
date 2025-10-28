import { Link } from 'react-router-dom';
import Navbar from '../components/Navbar';
import { Bell, Mail, MessageSquare, CheckCircle } from 'lucide-react';

function InboxPage() {
  return (
    <div className="min-h-screen bg-white">
      <Navbar />

      {/* Hero */}
      <section className="bg-gradient-to-br from-blue-500 to-blue-700 py-20">
        <div className="max-w-7xl mx-auto px-4 sm:px-6 lg:px-8">
          <div className="grid md:grid-cols-2 gap-12 items-center">
            <div>
              <h1 className="text-5xl font-bold text-white mb-6">Never miss an idea again</h1>
              <p className="text-xl text-blue-100 mb-8">
                FlowBoard Inbox captures everything - from emails to Slack messages - and turns them into actionable tasks.
              </p>
              <Link
                to="/register"
                className="inline-block bg-white text-blue-600 px-8 py-4 rounded-lg font-semibold hover:bg-gray-100 transition shadow-lg"
              >
                Try Inbox for free
              </Link>
            </div>
            <div className="bg-white rounded-2xl p-8 shadow-2xl">
              <h3 className="font-bold text-gray-900 mb-4">📬 Your Inbox</h3>
              <div className="space-y-3">
                <div className="bg-blue-50 p-4 rounded-lg border-l-4 border-blue-600">
                  <p className="font-medium text-gray-900">From: john@company.com</p>
                  <p className="text-sm text-gray-600">Review Q4 marketing plan</p>
                </div>
                <div className="bg-purple-50 p-4 rounded-lg border-l-4 border-purple-600">
                  <p className="font-medium text-gray-900">Slack: #design</p>
                  <p className="text-sm text-gray-600">Update landing page mockup</p>
                </div>
                <div className="bg-green-50 p-4 rounded-lg border-l-4 border-green-600 opacity-60">
                  <p className="font-medium text-gray-900 line-through">Quick idea</p>
                  <p className="text-sm text-gray-600">Blog post about automation</p>
                </div>
              </div>
            </div>
          </div>
        </div>
      </section>

      {/* Features */}
      <section className="py-20">
        <div className="max-w-7xl mx-auto px-4 sm:px-6 lg:px-8">
          <h2 className="text-4xl font-bold text-center text-gray-900 mb-16">How Inbox works</h2>
          <div className="grid md:grid-cols-3 gap-8">
            <div className="text-center">
              <div className="bg-blue-100 w-16 h-16 rounded-full flex items-center justify-center mx-auto mb-6">
                <Mail className="w-8 h-8 text-blue-600" />
              </div>
              <h3 className="text-xl font-bold text-gray-900 mb-4">Email Integration</h3>
              <p className="text-gray-600">Forward emails to your unique Inbox address. They automatically become tasks.</p>
            </div>
            <div className="text-center">
              <div className="bg-purple-100 w-16 h-16 rounded-full flex items-center justify-center mx-auto mb-6">
                <MessageSquare className="w-8 h-8 text-purple-600" />
                </div>
              <h3 className="text-xl font-bold text-gray-900 mb-4">Slack Integration</h3>
              <p className="text-gray-600">Save messages from Slack directly to your Inbox with one click.</p>
            </div>
            <div className="text-center">
              <div className="bg-green-100 w-16 h-16 rounded-full flex items-center justify-center mx-auto mb-6">
                <CheckCircle className="w-8 h-8 text-green-600" />
              </div>
              <h3 className="text-xl font-bold text-gray-900 mb-4">Organize & Act</h3>
              <p className="text-gray-600">Review, organize, and move tasks to the right board when ready.</p>
            </div>
          </div>
        </div>
      </section>

      {/* CTA */}
      <section className="bg-blue-50 py-20">
        <div className="max-w-4xl mx-auto px-4 text-center">
          <h2 className="text-4xl font-bold text-gray-900 mb-6">Start capturing ideas today</h2>
          <Link
            to="/register"
            className="inline-block bg-gradient-to-r from-blue-600 to-blue-700 text-white px-8 py-4 rounded-lg font-semibold hover:from-blue-700 hover:to-blue-800 transition shadow-lg"
          >
            Get started for free →
          </Link>
        </div>
      </section>
    </div>
  );
}

export default InboxPage;