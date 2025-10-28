import { Link } from 'react-router-dom';
import Navbar from '../components/Navbar';
import { Calendar, Clock, TrendingUp } from 'lucide-react';

function PlannerPage() {
  return (
    <div className="min-h-screen bg-white">
      <Navbar />

      {/* Hero */}
      <section className="bg-gradient-to-br from-purple-500 to-purple-700 py-20">
        <div className="max-w-7xl mx-auto px-4 sm:px-6 lg:px-8">
          <div className="grid md:grid-cols-2 gap-12 items-center">
            <div>
              <h1 className="text-5xl font-bold text-white mb-6">See your entire day at a glance</h1>
              <p className="text-xl text-purple-100 mb-8">
                FlowBoard Planner syncs with your calendar and helps you time-block your tasks for maximum productivity.
              </p>
              <Link
                to="/register"
                className="inline-block bg-white text-purple-600 px-8 py-4 rounded-lg font-semibold hover:bg-gray-100 transition shadow-lg"
              >
                Try Planner for free
              </Link>
            </div>
            <div className="bg-white rounded-2xl p-8 shadow-2xl">
              <h3 className="font-bold text-gray-900 mb-4">📅 Today's Schedule</h3>
              <div className="space-y-3">
                <div className="flex items-center gap-4">
                  <div className="text-sm font-medium text-gray-600 w-20">9:00 AM</div>
                  <div className="flex-1 bg-blue-100 p-3 rounded-lg border-l-4 border-blue-600">
                    <p className="font-medium text-gray-900">Team standup</p>
                  </div>
                </div>
                <div className="flex items-center gap-4">
                  <div className="text-sm font-medium text-gray-600 w-20">10:00 AM</div>
                  <div className="flex-1 bg-purple-100 p-3 rounded-lg border-l-4 border-purple-600">
                    <p className="font-medium text-gray-900">Deep work: Project review</p>
                  </div>
                </div>
                <div className="flex items-center gap-4">
                  <div className="text-sm font-medium text-gray-600 w-20">2:00 PM</div>
                  <div className="flex-1 bg-green-100 p-3 rounded-lg border-l-4 border-green-600">
                    <p className="font-medium text-gray-900">Client meeting</p>
                  </div>
                </div>
              </div>
            </div>
          </div>
        </div>
      </section>

      {/* Features */}
      <section className="py-20">
        <div className="max-w-7xl mx-auto px-4 sm:px-6 lg:px-8">
          <h2 className="text-4xl font-bold text-center text-gray-900 mb-16">Planner features</h2>
          <div className="grid md:grid-cols-3 gap-8">
            <div className="text-center">
              <div className="bg-purple-100 w-16 h-16 rounded-full flex items-center justify-center mx-auto mb-6">
                <Calendar className="w-8 h-8 text-purple-600" />
              </div>
              <h3 className="text-xl font-bold text-gray-900 mb-4">Calendar Sync</h3>
              <p className="text-gray-600">Connect Google Calendar, Outlook, or Apple Calendar. See everything in one place.</p>
            </div>
            <div className="text-center">
              <div className="bg-blue-100 w-16 h-16 rounded-full flex items-center justify-center mx-auto mb-6">
                <Clock className="w-8 h-8 text-blue-600" />
              </div>
              <h3 className="text-xl font-bold text-gray-900 mb-4">Time Blocking</h3>
              <p className="text-gray-600">Drag and drop tasks into your calendar. Allocate time for deep work.</p>
            </div>
            <div className="text-center">
              <div className="bg-green-100 w-16 h-16 rounded-full flex items-center justify-center mx-auto mb-6">
                <TrendingUp className="w-8 h-8 text-green-600" />
              </div>
              <h3 className="text-xl font-bold text-gray-900 mb-4">Weekly Views</h3>
              <p className="text-gray-600">Plan your entire week. See upcoming deadlines and meetings.</p>
            </div>
          </div>
        </div>
      </section>

      {/* CTA */}
      <section className="bg-purple-50 py-20">
        <div className="max-w-4xl mx-auto px-4 text-center">
          <h2 className="text-4xl font-bold text-gray-900 mb-6">Take control of your time</h2>
          <Link
            to="/register"
            className="inline-block bg-gradient-to-r from-purple-600 to-purple-700 text-white px-8 py-4 rounded-lg font-semibold hover:from-purple-700 hover:to-purple-800 transition shadow-lg"
          >
            Get started for free →
          </Link>
        </div>
      </section>
    </div>
  );
}

export default PlannerPage;