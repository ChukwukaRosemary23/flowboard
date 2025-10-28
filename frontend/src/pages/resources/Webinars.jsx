import { Link } from 'react-router-dom';
import { ArrowLeft, Video, Calendar, Users } from 'lucide-react';
import Navbar from '../../components/Navbar';

export default function Webinars() {
  const webinars = [
    {
      id: 1,
      title: "FlowBoard Fundamentals",
      date: "Nov 5, 2024 at 2:00 PM",
      duration: "1 hour",
      attendees: "250+ registered",
      status: "Upcoming",
      image: "https://images.unsplash.com/photo-1587825140708-dfaf72ae4b04?w=800"
    },
    {
      id: 2,
      title: "Advanced Automation Techniques",
      date: "Nov 12, 2024 at 3:00 PM",
      duration: "1.5 hours",
      attendees: "180+ registered",
      status: "Upcoming",
      image: "https://images.unsplash.com/photo-1531482615713-2afd69097998?w=800"
    },
    {
      id: 3,
      title: "Building High-Performance Teams",
      date: "Oct 20, 2024",
      duration: "1 hour",
      attendees: "500+ watched",
      status: "Recorded",
      image: "https://images.unsplash.com/photo-1552664730-d307ca884978?w=800"
    }
  ];

  return (
    <div className="min-h-screen bg-gray-50">
      <Navbar />
      
      <div className="max-w-7xl mx-auto px-4 sm:px-6 lg:px-8 py-12">
        <div className="mb-8">
          <Link to="/" className="inline-flex items-center text-blue-600 hover:text-blue-700 mb-4">
            <ArrowLeft className="w-4 h-4 mr-2" />
            Back to Home
          </Link>
          <h1 className="text-4xl font-bold text-gray-900 mb-4">FlowBoard Webinars</h1>
          <p className="text-xl text-gray-600">
            Join live sessions or watch recordings to level up your skills
          </p>
        </div>

        <div className="space-y-6">
          {webinars.map((webinar) => (
            <div key={webinar.id} className="bg-white rounded-lg shadow-md overflow-hidden hover:shadow-xl transition">
              <div className="md:flex">
                <div className="md:w-1/3">
                  <img src={webinar.image} alt={webinar.title} className="w-full h-full object-cover" />
                </div>
                <div className="p-6 md:w-2/3">
                  <div className="flex items-center justify-between mb-4">
                    <span className={`px-3 py-1 rounded-full text-sm font-medium ${
                      webinar.status === 'Upcoming' 
                        ? 'bg-green-100 text-green-700' 
                        : 'bg-gray-100 text-gray-700'
                    }`}>
                      {webinar.status}
                    </span>
                    <Video className="w-6 h-6 text-blue-600" />
                  </div>
                  <h2 className="text-2xl font-bold text-gray-900 mb-3">{webinar.title}</h2>
                  <div className="space-y-2 text-gray-600 mb-4">
                    <div className="flex items-center">
                      <Calendar className="w-4 h-4 mr-2" />
                      <span>{webinar.date}</span>
                    </div>
                    <div className="flex items-center">
                      <Clock className="w-4 h-4 mr-2" />
                      <span>{webinar.duration}</span>
                    </div>
                    <div className="flex items-center">
                      <Users className="w-4 h-4 mr-2" />
                      <span>{webinar.attendees}</span>
                    </div>
                  </div>
                  <button className="bg-blue-600 text-white px-6 py-2 rounded-lg hover:bg-blue-700 transition">
                    {webinar.status === 'Upcoming' ? 'Register Now' : 'Watch Recording'}
                  </button>
                </div>
              </div>
            </div>
          ))}
        </div>
      </div>
    </div>
  );
}