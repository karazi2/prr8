import 'package:dio/dio.dart';
import 'note.dart';

class ApiService {
  final Dio _dio = Dio();

  Future<List<Note>> getApartments() async {
    try {
      final response = await _dio.get('http://192.168.0.19:8080/apartments');
      if (response.statusCode == 200) {
        List<Note> apartments = (response.data as List)
            .map((apartment) => Note.fromJson(apartment))
            .toList();
        return apartments;
      } else {
        throw Exception('Failed to load apartments');
      }
    } catch (e) {
      throw Exception('Error fetching apartments: $e');
    }
  }
}
